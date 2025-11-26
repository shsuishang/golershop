package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"net/http"
	"time"

	"github.com/gogf/gf/contrib/config/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// initNacosConfig 使用Nacos 配置信息
func initNacosConfig(ctx context.Context) error {
	// Get Nacos configuration from config file
	nacosIP := g.Cfg().MustGet(ctx, "nacos.server.ip").String()
	nacosPort := g.Cfg().MustGet(ctx, "nacos.server.port").Uint64()

	// Get Nacos client configuration
	cacheDir := g.Cfg().MustGet(ctx, "nacos.client.cacheDir").String()
	logDir := g.Cfg().MustGet(ctx, "nacos.client.logDir").String()

	timeoutMs := g.Cfg().MustGet(ctx, "nacos.client.timeoutMs").Uint64()
	if timeoutMs == 0 {
		timeoutMs = 5000 // Default value
	}

	// Get Nacos config parameters
	dataId := g.Cfg().MustGet(ctx, "nacos.config.dataId").String()
	if dataId == "" {
		dataId = "golershop.yaml" // Default value
	}

	group := g.Cfg().MustGet(ctx, "nacos.config.group").String()
	if group == "" {
		group = "GOLERSHOP_GROUP" // Default value
	}

	// Configure Nacos server connection
	// This defines how to connect to the Nacos server
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacosIP,   // Nacos server address
			Port:   nacosPort, // Nacos server port
		},
	}

	// Configure Nacos client settings
	// This defines local cache and logging behavior
	clientConfig := constant.ClientConfig{
		CacheDir:  cacheDir,  // Directory for local cache
		LogDir:    logDir,    // Directory for log files
		TimeoutMs: timeoutMs, // Set timeout
	}

	// Configure configuration parameters
	// This defines which configuration to retrieve
	configParam := vo.ConfigParam{
		DataId: dataId, // Configuration file identifier
		Group:  group,  // Configuration group name
	}

	// Check if Nacos server is accessible before trying to create adapter
	// Try multiple health check endpoints for different Nacos versions
	if !isNacosServerAccessible(nacosIP, nacosPort, time.Duration(timeoutMs)*time.Millisecond) {
		err := gerror.Newf("Nacos server %s:%d is not accessible", nacosIP, nacosPort)
		glog.Warningf(ctx, "Nacos server %s:%d is not accessible, falling back to local configuration", nacosIP, nacosPort)
		panic(err)

		return err // Return error to indicate Nacos server is not accessible
	}

	if err := initNacosRegistry(ctx); err != nil {
	}

	// Try to create Nacos adapter with configuration
	// The adapter implements gcfg.Adapter interface for configuration management
	adapter, err := nacos.New(ctx, nacos.Config{
		ServerConfigs: serverConfigs, // Server connection settings
		ClientConfig:  clientConfig,  // Client behavior settings
		ConfigParam:   configParam,   // Configuration retrieval settings
	})

	// Instead of returning fatal error, we just log it and continue with local config
	if err != nil {
		glog.Warningf(ctx, "Failed to create Nacos adapter: %+v", err)
		return err // Return error to indicate failure to create Nacos adapter
	}

	// Set Nacos adapter as the configuration adapter
	// This enables GoFrame to use Nacos for configuration management
	g.Cfg().SetAdapter(adapter)

	return nil
}

// isNacosServerAccessible checks if the Nacos server is accessible
// Try multiple health check endpoints for different Nacos versions
func isNacosServerAccessible(ip string, port uint64, timeout time.Duration) bool {
	// List of possible health check endpoints for different Nacos versions
	healthEndpoints := []string{
		"/nacos/v1/console/health", // Older versions
		//"/nacos/actuator/health",   // Newer versions (Spring Boot Actuator)
		//"/nacos/health",            // Some versions
		//"/v1/console/health",       // Relative path
		//"/actuator/health",         // Spring Boot Actuator relative path
	}

	serverAddr := fmt.Sprintf("%s:%d", ip, port)

	// Try each endpoint until one works
	for _, endpoint := range healthEndpoints {
		healthURL := "http://" + serverAddr + endpoint

		client := &http.Client{
			Timeout: timeout,
		}

		resp, err := client.Get(healthURL)
		if err != nil {
			continue // Try next endpoint
		}

		// Close response body
		resp.Body.Close()

		// If we get a successful response (2xx), the server is accessible
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			glog.Debugf(context.Background(), "Nacos health check successful with endpoint: %s", endpoint)
			return true
		}

		// For Nacos 3.1.0, even 404 might indicate the server is running
		// (some endpoints may not exist but the server is still accessible)
		if resp.StatusCode == http.StatusNotFound {
			// Try a simple connection test
			testURL := "http://" + serverAddr + "/nacos"
			testResp, testErr := client.Get(testURL)
			if testErr == nil {
				testResp.Body.Close()
				// If we can reach the main page, consider it accessible
				if testResp.StatusCode >= 200 && testResp.StatusCode < 500 {
					glog.Debugf(context.Background(), "Nacos server accessible (main page), endpoint: %s", endpoint)
					return true
				}
			}
		}
	}

	// If none of the endpoints work, try a basic connectivity test
	// Just try to connect to the server without a specific endpoint
	connURL := "http://" + serverAddr
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(connURL)
	if err != nil {
		glog.Debugf(context.Background(), "Nacos server not accessible: %v", err)
		return false
	}
	defer resp.Body.Close()

	// If we can connect to the server at all, consider it accessible
	glog.Debugf(context.Background(), "Nacos server accessible via basic connection test")
	return resp.StatusCode < 500
}
