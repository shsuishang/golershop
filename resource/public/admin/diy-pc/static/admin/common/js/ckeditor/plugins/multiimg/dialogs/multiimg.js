(function() {
    var $w = $(window).width() * 0.8 + 'px';
    var $h = $(window).height() * 0.8 + 'px';

    var rid = $.cookie('rid');
    var img_url = '';

    if (2 == rid)
    {
        img_url = SYS.CONFIG.index_page + '?ctl=Store_Media&met=index&typ=e&mdu=seller&opener=ckeditor&v=' +new Date().getSeconds();
    }
    else
    {
        img_url = SYS.CONFIG.index_page + '?ctl=Store_Media&met=index&typ=e&opener=ckeditor&v=' +new Date().getSeconds();
    }

    CKEDITOR.dialog.add("multiimg",
                        function(a) {
                            var ROOT_PATH = "/study/"; // your root path

                            return {
                                title: __("批量上传图片"),
                                minWidth: "660px",
                                minHeight:"400px",
                                contents: [{
                                    id: "tab1",
                                    label: "",
                                    title: "",
                                    expand: true,
                                    width: $w,
                                    height: $h,
                                    padding: 0,
                                    padding: 0,
                                    elements: [{
                                        type: "html",
                                        style: "width:" + $w + ";height:" + $h,
                                        html: '<iframe id="uploadFrame" src="' + img_url + '" frameborder="0"></iframe>'
                                    }]
                                }],
                                // when the dialog ended width ensure,"onOK" will be executed.
                                onOk: function() {
                                    var ins = a;
                                    var num = window.selectImgList.length;
                                    var imgHtml = "";
                                    for(var i=0;i<num;i++){
                                        imgHtml += "<p><img src=\"" + window.selectImgList[i] + "\" /></p>";
                                    }
                                    ins.insertHtml(imgHtml);
                                },
                                onShow: function () {
                                    document.getElementById("uploadFrame").setAttribute("src", img_url);
                                }
                            }
                        })
})();
