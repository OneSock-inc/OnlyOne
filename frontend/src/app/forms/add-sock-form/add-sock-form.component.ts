import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { DataUrl, NgxImageCompressService } from 'ngx-image-compress';

@Component({
  selector: 'app-add-sock-form',
  templateUrl: './add-sock-form.component.html',
  styleUrls: ['./add-sock-form.component.scss']
})
export class AddSockFormComponent implements OnInit {

  constructor(private imageCompress: NgxImageCompressService) { }


  addSockForm!: FormGroup;

  pictureButtonLabel = "Take picture";

  sizeValue: number = 36;

  pictureB64: DataUrl = "";
  
  maxPxBorder: number = 500;

  sockColor: string = "#ffffff";
  colorPickerLabel: string = "Choose color";

  screenWidth!: string;
  textColor: string = "#ffffff";
  
  ngOnInit(): void {
    this.addSockForm = new FormGroup({
      description: new FormControl('',{
        validators: [Validators.required]
      }),
      picture: new FormControl('',{
        validators: [Validators.required]
      }),
      shoeSize: new FormControl('',{
        validators: [Validators.required]
      }),
      sockType: new FormControl('',{
        validators: [Validators.required]
      }),
    });
    this.screenWidth = this.getScreenWidth().toString();
  }

  onSubmit(form: FormGroup): void {
    alert("Sock added successfully");
    // send to api
    //pictureB64
    //sockColor
    //form.description
    //form.password

  }

  getScreenWidth(): number {
    if (window.innerWidth < 500) {
      let formWidth = document.getElementsByTagName("form")[0].clientWidth;
      return Math.floor(formWidth * 0.8); // TODO : find a way to get the width of the form
    }
    return 500;
  }

  onColorChange(newColor: string): void {
    this.colorPickerLabel = "Change color";
    let colorShower = document.getElementById("colorShower");
    if (colorShower) {
      colorShower.style.backgroundColor = newColor;
    }
    
  }


}
