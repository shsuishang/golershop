$(function() {
    $.ajaxSetup({
      headers : {
        'Authorization' : header_auth_token
      }
    });
    var $grid = $("#grid");
    var $form = $("#manage-form");
    var $handle = $.extend(handle, {
        initDom: function() {

            if ($('#category_id').length > 0)
            {
                //商品类别
                var opts = {
                    url: api_url + "/front/pt/product/listAllCategory",
                    width : 300,
                    selectOnlyLeaf : false,
                    //inputWidth : (SYSTEM.enableStorage ? 145 : 208),
                    inputWidth :  300,
                    //defaultSelectValue : '-1',
                    //defaultSelectValue : rowData.categoryId || '',
                    showRoot : true
                }

                var categoryTree = Public.categoryTree($('#category_id'), opts, 'product_category');
            }

            return this;
        },

        initEvent: function() {
            initSelect2Brand('brand_id');

            return this;
        },

        initField: function(rowData) {
            if (rowData.id) {
                $('#item_id').val(rowData.item_id);
                $('#item_name').val(rowData.item_name);
                $('#product_id').val(rowData.product_id);
                $('#color_id').val(rowData.color_id);
                $('#item_is_default').val(rowData.item_is_default);
                $('#item_number').val(rowData.item_number);
                $('#item_barcode').val(rowData.item_barcode);
                $('#item_cost_price').val(rowData.item_cost_price);
                $('#item_unit_price').val(rowData.item_unit_price);
                $('#item_market_price').val(rowData.item_market_price);
                $('#item_quantity').val(rowData.item_quantity);
                $('#item_warn_quantity').val(rowData.item_warn_quantity);
                $('#item_spec').val(rowData.item_spec);
                $('#item_enable').val(rowData.item_enable);
                $('#item_is_change').bootstrapSwitch('state', rowData.item_is_change);

                //$('#' + this.$priKey).attr("readonly", "readonly");
                //$('#' + this.$priKey).addClass('ui-input-dis');
                this.initState();
            }

            return this;
        },

        resetForm: function(t) {
            $('#item_id').val('');
            $('#item_name').val('');
            $('#product_id').val('');
            $('#color_id').val('');
            $('#item_is_default').val('');
            $('#item_number').val('');
            $('#item_barcode').val('');
            $('#item_cost_price').val('');
            $('#item_unit_price').val('');
            $('#item_market_price').val('');
            $('#item_quantity').val('');
            $('#item_warn_quantity').val('');
            $('#item_spec').val('');
            $('#item_enable').val('');
            $('#item_is_change').bootstrapSwitch('state', 0);

            this.initState();

            return this;
        }
    });

    var $col_model = [{
        "name": "operate",
        "label": __('操作'),
        "width": 80,
        "sortable": false,
        "search": false,
        "resizable": false,
        "frozen": false,
        "fixed": true,
        "hidden": true,
        "align": "center",
        "title": true,
        "formatter": $handle.operFormatter
    }, {
        "name": "item_id",
        "index": "item_id",
        "label": __('SKU编号'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "width": 100
    }, {
        "name": "product_name",
        "index": "product_name",
        "label": __('商品名称'),
        "classes": "ui-ellipsis",
        "align": "left",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "item_spec_name",
        "index": "item_spec_name",
        "label": __('商品名称'),
        "classes": "ui-ellipsis",
        "align": "left",
        "title": true,
        "width": 100
    }, {
        "name": "item_name",
        "index": "item_name",
        "label": __('商品名称'),
        "classes": "ui-ellipsis",
        "align": "left",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "product_id",
        "index": "product_id",
        "label": __('产品SPU'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "color_id",
        "index": "color_id",
        "label": __('颜色SKU'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "hidden": true,
        "width": 60
    }, {
        "name": "item_is_default",
        "index": "item_is_default",
        "label": __('是否默认'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "hidden": true,
        "width": 60
    }, {
        "name": "item_number",
        "index": "item_number",
        "label": "SKU商家编码",
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "item_barcode",
        "index": "item_barcode",
        "label": __('条形码'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "item_cost_price",
        "index": "item_cost_price",
        "label": __('成本价'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        hidden: SYS.STORE_SUPPLIER_ENABLE ? false : true,
        "width": 100
    }, {
        "name": "item_platform_price",
        "index": "item_platform_price",
        "label": __('平台价'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        hidden: SYS.STORE_SUPPLIER_ENABLE ? false : true,
        editable: SYS.STORE_SUPPLIER_ENABLE ? true : false,
        "width": 100
    }, {
        "name": "item_unit_price",
        "index": "item_unit_price",
        "label": __('商品价格'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        hidden: false,
        editable: SYS.STORE_SUPPLIER_ENABLE ? true : false,
        "width": 100
    }, {
        "name": "item_market_price",
        "index": "item_market_price",
        "label": __('市场价'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        hidden: SYS.STORE_SUPPLIER_ENABLE ? false : true,
        editable: SYS.STORE_SUPPLIER_ENABLE ? true : false,
        "width": 100
    }, {
        "name": "item_quantity",
        "index": "item_quantity",
        "label": __('商品库存'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "width": 60
    }, {
        "name": "product_image",
        "index": "product_image",
        "label": __('商品图片'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "hidden": true,
        "width": 60
    },{
        "name": "item_spec",
        "index": "item_spec",
        "label": __('商品规格序列化'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "hidden": true,
        "width": 100
    }, {
        "name": "item_enable",
        "index": "item_enable",
        "label": __('是否启用'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "width": 60,
        "formatter": function(val, opt, row) {
            var r = {
                "1001": __('销售中'),
                "1002": __('仓库中'),
                "1000": __('违规禁售')
            };
            return r[val];
        }
    }, {
        "name": "item_is_change",
        "index": "item_is_change",
        "label": __('被改动'),
        "classes": "ui-ellipsis",
        "align": "center",
        "title": true,
        "fixed": true,
        "width": 60,
        "hidden": true,
        "formatter": function(val, opt, row) {
            var r = [
                __('未改动'),
                __('已改动分销使用')
            ];
            return r[val];
        }
    }];

    $handle.init($grid, $form, 'item_id', 'Product_Item', $.getUrlParam('mdu'));

    $handle.$url.lists = api_url + "/front/pt/product/listItem";

/*
    if ($.getUrlParam('mdu'))
    {
        $handle.$url.lists = $handle.$url.lists + "&mdu=" + $.getUrlParam('mdu')
    }*/


    var $opt = {
        method:"GET",
        cellEdit: true,
        cellsubmit : 'remote',
        cellurl : SYS.CONFIG.index_url+'?ctl=Product_Item&met=editFiled&typ=json&product_id=' + $.getUrlParam('product_id')
    };


    //manage
    if (frameElement && frameElement.api) {
        //var curRow, curCol, curArrears;
        var api = frameElement.api;

        if (api.data.isFilter)
        {
            $opt = {
                multiselect: true
            };

            $('#search_box').show()
        }
        else
        {
            $('#search_box').hide()
        }

        if (typeof api.data.selectList=='string')
        {
            var ids = api.data.selectList;
            var id_arr = ids.split(',');
            for ( var i = 0, len = id_arr.length; i < len; i++){
                var rowid = id_arr[i];
                var rowData = $grid.jqGrid('getRowData', rowid);
                selectList[rowid] = rowData;
            }
        }
    }


    if ($grid.length > 0) {
        $handle.initDom().initGrid($col_model, $opt).initGridEvent();
    }

    //manage
    if (frameElement && frameElement.api) {
        //var curRow, curCol, curArrears;
        var api = frameElement.api;


        api['config']['ok'] = true;
        $handle.initPopBtns(api, {
            fields: {},
        });

        $handle.initField(api.data.rowData || {}).initState();
    }
});
