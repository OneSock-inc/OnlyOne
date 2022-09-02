import { Component, OnInit, HostListener } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import {DataUrl, NgxImageCompressService} from "ngx-image-compress";


@Component({
  selector: 'app-add-sock-page',
  templateUrl: './add-sock-page.component.html',
  styleUrls: ['./add-sock-page.component.scss'],
  host: {'class': 'default-layout'}
})
export class AddSockPageComponent implements OnInit {

  displayArrow: boolean = true;

  addSockForm!: FormGroup;

  pictureButtonLabel = "Take picture";

  sizeValue: number = 36;

  pictureB64: DataUrl = "";
  
  maxPxBorder: number = 500;

  sockColor: string = "#ffffff";
  colorPickerLabel: string = "Choose color";

  screenWidth!: string;
  textColor: string = "#ffffff";
  


  constructor(private imageCompress: NgxImageCompressService) { }

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

  onSubmit(form: any): void {
    alert("Sock added successfully");
    // send to api
    //pictureB64
    //sockColor
    //form.description
    //form.password

  }

  // display down arrow if the user has not scrolled to the bottom of the page
  @HostListener('window:scroll', ['$event'])
  onScroll(event: Event): void {
    if (window.pageYOffset >= (document.documentElement.scrollHeight - document.documentElement.clientHeight)) {
      this.displayArrow = false;
    }
    else {
      this.displayArrow = true;
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
  
  compressFile(image: DataUrl) {   
    this.imageCompress.compressFile(image, 1, 100, 100, this.maxPxBorder, this.maxPxBorder).then(
      result => {
        this.pictureB64 = result;
      }
    );
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
