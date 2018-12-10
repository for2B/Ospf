import { Component, OnInit ,Input,Output,EventEmitter} from '@angular/core';
import { SITE_HOST_URL,UPLOAD_FILE } from '../../const/const_chenhuiliang'
@Component({
  selector: 'froala-editor',
  templateUrl: './froala-editor.component.html',
  styleUrls: ['./froala-editor.component.css']
})
export class FroalaEditorComponent implements OnInit {
  froalaText:string;
  @Output() froala = new EventEmitter<string>();
  @Input() initialValues :string
  option:Object;
  toolbarbuttons=[
   'fullscreen','bold', 'italic', 'underline', 'strikeThrough',
   '|', 'fontFamily', 'fontSize', 'color',
   '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent','insertLink', 'insertImage', 'embedly', 'insertFile', 'insertTable',
   '|','specialCharacters', 'insertHR', 'selectAll', 'clearFormatting',
   '|',  'spellChecker', '|', 'undo', 'redo'
  ];

  constructor() {
    this.froalaText = "";
  }

  ngOnInit() {
    if(this.initialValues){
      this.froalaText = this.initialValues
    }
    var that = this
    this.option = {
      language:"zh_cn",
      placeholederText:"请输入内容",
      charCounterCount:true,
      charCounterMax:400,
      toolbarButtons:this.toolbarbuttons,
      toolbarButtonsSM:this.toolbarbuttons,
      toolbarButtonsXS:this.toolbarbuttons,
      codeMirror:false,
      codeMirrorOptions:{
        tabSize:4
      },
      //配置文件上传
      fileUploadParam:'filename',
      fileUploadURL:SITE_HOST_URL+UPLOAD_FILE,
      fileUploadMethod:'POST',
      //配置图片上传
      imageUploadParam:'filename',
      imageUploadMethod:'POST',
      imageUploadURL:SITE_HOST_URL+UPLOAD_FILE,
      //size
      height:300,
      // width:800,
      events:{
        'froalaEditor.keyup':function(e,editor){
          that.froala.emit(that.froalaText);
        },
        'froalaEditor.blur':function(e,editor,touchendEvent){
          that.froala.emit(that.froalaText);
        },
        'froalaEditor.file.error':function(e,editor,error,response){
          if (error.code == 1) {
            console.log("Bad link")
           }

          // No link in upload response.
          else if (error.code == 2) {
            console.log(" No link in upload response.")

           }

          // Error during file upload.
          else if (error.code == 3) {
            console.log("Error during file upload.")

           }

          // Parsing response failed.
          else if (error.code == 4) {
            console.log("Parsing response failed.")

           }

          // File too text-large.
          else if (error.code == 5) {
            console.log("File too text-large.")

           }

          // Invalid file type.
          else if (error.code == 6) {
            console.log("Invalid file type")

           }

          // File can be uploaded only to same domain in IE 8 and IE 9.
          else if (error.code == 7) {
            console.log(" File can be uploaded only to same domain in IE 8 and IE 9.")

           }

          // Response contains the original server response to the request if available.
        }
      }
    };
  }
}
