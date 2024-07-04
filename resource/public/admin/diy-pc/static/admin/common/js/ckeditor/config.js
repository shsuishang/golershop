/**
 * @license Copyright (c) 2003-2020, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see https://ckeditor.com/legal/ckeditor-oss-license
 */

CKEDITOR.config.plugins='dialogui,dialog,about,a11yhelp,colorbutton,preview,templates,widget,basicstyles,blockquote,clipboard,panel,floatpanel,menu,contextmenu,resize,button,toolbar,elementspath,enterkey,entities,popup,filebrowser,floatingspace,font,justify,listblock,richcombo,format,horizontalrule,htmlwriter,wysiwygarea,image,multiimg,indent,indentlist,fakeobjects,link,list,magicline,maximize,pastetext,pastefromword,removeformat,showborders,sourcearea,specialchar,menubutton,scayt,stylescombo,tab,table,tabletools,undo,wsc';

CKEDITOR.editorConfig = function( config ) {   
	// Define changes to default configuration here. For example:
	// config.language = 'fr';
	// config.uiColor = '#AADC6E';
	// Define changes to default configuration here.
	// For complete reference see:
	// http://docs.ckeditor.com/#!/api/CKEDITOR.config
    
/*    config.toolbar =
        [
            [ 'Bold', 'Italic', '-', 'NumberedList', 'BulletedList', '-', 'Link', 'Unlink','-','multiimg' ]
        ]
	*/
	// The toolbar groups arrangement, optimized for two toolbar rows.
	config.toolbarGroups = [
		{ name: 'clipboard',   groups: [ 'clipboard', 'undo' ] },
		{ name: 'editing',     groups: [ 'find', 'selection', 'spellchecker' ] },
		{ name: 'links' },
        { name: 'others' },
		{ name: 'insert' },
		{ name: 'forms' },
		{ name: 'tools' },
		{ name: 'basicstyles', groups: [ 'basicstyles', 'cleanup' ] },
		{ name: 'paragraph',   groups: [ 'list', 'indent', 'blocks', 'align', 'bidi' ] },
		{ name: 'styles' },
		{ name: 'colors' },
		{ name: 'templates' },
		{ name: 'widget' },
		{ name: 'document',	   groups: [ 'mode', 'document', 'doctools' ] },
		{ name: 'preview' }
	];

	// Remove some buttons provided by the standard plugins, which are
	// not needed in the Standard(s) toolbar.
	config.removeButtons = 'Underline,Subscript,Superscript';

	// Set the most common block elements.
	//config.format_tags = 'p;h1;h2;h3;pre';

    config.filebrowserImageBrowseUrl = 'admin.php?ctl=Store_Media&met=index&typ=e&mdu=seller';
    // config.filebrowserBrowseUrl = 'kcfinder/browse.php?type=files';
    // config.filebrowserImageBrowseUrl = 'kcfinder/browse.php?type=images';
    // config.filebrowserFlashBrowseUrl = 'kcfinder/browse.php?type=flash';
    // config.filebrowserUploadUrl = 'kcfinder/upload.php?type=files';
    // config.filebrowserImageUploadUrl = 'kcfinder/upload.php?type=images';
    // config.filebrowserFlashUploadUrl = 'kcfinder/upload.php?type=flash';

    // Simplify the dialog windows.
	config.removeDialogTabs = 'image:advanced;link:advanced';
	
	// BootstrapCK Skin Options
	config.skin = 'bootstrapck';
	config.height = '350px';
};


$(function () {

	/*
	//需要手动更新CKEDITOR字段
    for ( instance in CKEDITOR.instances )
	{
        // When the CKEDITOR instance is created, fully initialized and ready for interaction.
		// 当id为content的那个ckeditor被创建，并初始化完成之后
        CKEDITOR.instances[instance].on("instanceReady", function() {
            // 保存按钮
            this.addCommand("image", {
                modes : {
                    wysiwyg : 1,
                    source : 1
                },
                exec : function(editor) {
                    // 获取到editor中的内容
                    var content = editor.document.getBody().getHtml();
                    alert(content);
                }
            });
        });
	}
 	*/
	
})


