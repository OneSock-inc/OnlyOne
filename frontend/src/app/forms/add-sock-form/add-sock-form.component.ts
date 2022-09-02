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
  }

  onSubmit(form: FormGroup): void {
    alert("Sock added successfully");
    // send to api
    //pictureB64
    //sockColor
    //form.description
    //form.password

  }

  onColorChange(newColor: string): void {
    this.colorPickerLabel = "Change color";
    let colorShower = document.getElementById("colorShower");
    if (colorShower) {
      colorShower.style.backgroundColor = newColor;
    }
    
  }

  selectFile(event: any) {
    if (event.target.files && event.target.files[0]) {
      let reader = new FileReader();
      reader.onload = (event: any) => {
        this.compressFile(event.target.result)
      }
      reader.readAsDataURL(event.target.files[0]);
      this.pictureButtonLabel = "Change picture";
    }
    else {
      this.pictureButtonLabel = "Take picture";
    }
    
  }

  private compressFile(image: DataUrl) {   
    this.imageCompress.compressFile(image, 1, 100, 100, this.maxPxBorder, this.maxPxBorder).then(
      result => {
        this.pictureB64 = result;
      }
    );
  }


}
