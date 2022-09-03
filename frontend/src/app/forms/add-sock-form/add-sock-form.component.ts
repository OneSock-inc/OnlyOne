import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { DataUrl, NgxImageCompressService } from 'ngx-image-compress';

import { Sock, SockType} from 'src/app/dataModel/sock.model';

@Component({
  selector: 'app-add-sock-form',
  templateUrl: './add-sock-form.component.html',
  styleUrls: ['./add-sock-form.component.scss']
})
export class AddSockFormComponent implements OnInit {

  constructor(private imageCompress: NgxImageCompressService, private http: HttpClient) { 
    this.newSock = new Sock();
  }

  newSock: Sock;

  // Sepcial fields values
  // Init with thtis.initForm() function
  sizeValue!: number;
  pictureB64!: DataUrl;
  sockColor!: string;

  // Sepcial fields values
  pictureButtonLabel = "Take picture";
  colorPickerLabel: string = "Choose color";
  
  // Forms model
  addSockForm!: FormGroup;

  // Aspect
  maxPxBorder: number = 500;
  textColor: string = "#ffffff";
  screenWidth!: string;

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
        validators: [Validators.required, Validators.min(SockType.low), Validators.max(SockType.high)]
      }),
    });
    this.initForm();
    this.screenWidth = this.getScreenWidth().toString()
  }

  onSubmit(form: FormGroup): void {
    if (!form.valid) return;
    this.newSock.shoeSize = form.value.shoeSize;
    this.newSock.color = this.sockColor;
    this.newSock.description = form.value.description;
    this.newSock.type = Number(form.value.sockType);
    this.newSock.picture = this.pictureB64.split(',')[1];

    const newSockStr = this.newSockToJson(this.newSock);

    this.http.post<any>("https://api.jsch.ch/sock/", newSockStr)
      .subscribe({
        next: data => {
          console.log(data);
          this.initForm();
          alert('New sock successfully added !');
        },
        error: err => {
          console.log(err);
          alert(err.message)
        }
      })
    
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

  private initForm() {
    this.sizeValue = 36;
    this.pictureB64 = '';
    this.sockColor = "#ffffff";
  }

  private getScreenWidth(): number {
    if (window.innerWidth < 500) {
      let windowWidth = window.screen.width;
      //let formWidth = document.getElementsByTagName("form")[0].clientWidth;
      return Math.floor(windowWidth * 0.8); // TODO : find a way to get the width of the form
    }
    return 500;
  }

  private newSockToJson(newSock: Sock): string {
    return JSON.stringify(this.newSock, (key, value) => {
      if (value === ''){
        return undefined;
      } else {
        return value;
      }
    });
  }
  
}
