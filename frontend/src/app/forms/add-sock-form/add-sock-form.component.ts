import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { DataUrl, NgxImageCompressService } from 'ngx-image-compress';

import { Sock, SockType } from 'src/app/dataModel/sock.model';
import {
  PostResponse,
  SocksManagerService,
} from 'src/app/services/socksManager/socks-manager.service';

@Component({
  selector: 'app-add-sock-form',
  templateUrl: './add-sock-form.component.html',
  styleUrls: ['./add-sock-form.component.scss'],
})
export class AddSockFormComponent implements OnInit {
  constructor(
    private imageCompress: NgxImageCompressService,
    private socksMan: SocksManagerService,
    private router: Router
  ) {
    this.newSock = new Sock();
  }

  private newSock: Sock;

  // Sepcial fields values
  // Init with thtis.initForm() function
  sizeValue!: number;
  pictureB64!: DataUrl;
  sockColor!: string;

  // Sepcial fields values
  pictureButtonLabel = 'Take picture';
  colorPickerLabel: string = 'Choose color';

  // Forms model
  addSockForm!: FormGroup;

  // Aspect
  maxPxBorder: number = 500;
  textColor: string = '#ffffff';
  screenWidth!: string;

  ngOnInit(): void {
    this.addSockForm = new FormGroup({
      description: new FormControl('', {
        validators: [Validators.required],
      }),
      picture: new FormControl('', {
        validators: [Validators.required],
      }),
      shoeSize: new FormControl('', {
        validators: [Validators.required],
      }),
      sockType: new FormControl('', {
        validators: [
          Validators.required,
          Validators.min(SockType.low),
          Validators.max(SockType.high),
        ],
      }),
    });
    this.initForm();
    this.screenWidth = this.getScreenWidth().toString();
  }

  onSubmit(form: FormGroup): void {
    if (!form.valid) return;
    form.disable();
    this.newSock.shoeSize = form.value.shoeSize;
    this.newSock.color = this.sockColor;
    this.newSock.description = form.value.description;
    this.newSock.type = Number(SockType[form.value.sockType]);
    this.newSock.picture = this.pictureB64.split(',')[1];

    this.socksMan.registerNewSock(this.newSock).subscribe({
      next: (response: PostResponse) => {
        alert(`New sock added !)`);
        this.router.navigate(['/home']);
      },
      error: (e) => alert(`ERROR : ${e.message}`),
    });
  }

  onColorChange(newColor: string): void {
    this.colorPickerLabel = 'Change color';
    let colorShower = document.getElementById('colorShower');
    if (colorShower) {
      colorShower.style.backgroundColor = newColor;
    }
  }

  selectFile(event: any) {
    if (event.target.files && event.target.files[0]) {
      let reader = new FileReader();
      reader.onload = (event: any) => {
        this.compressFile(event.target.result);
      };
      reader.readAsDataURL(event.target.files[0]);
      this.pictureButtonLabel = 'Change picture';
    } else {
      this.pictureButtonLabel = 'Take picture';
    }
  }

  private compressFile(image: DataUrl) {
    this.imageCompress
      .compressFile(image, 1, 100, 100, this.maxPxBorder, this.maxPxBorder)
      .then((result) => {
        this.pictureB64 = result;
      });
  }

  private initForm() {
    this.addSockForm.reset();
    this.sizeValue = 40;
    this.pictureB64 = '';
    this.sockColor = '#ffffff';
  }

  private getScreenWidth(): number {
    if (window.innerWidth < 500) {
      let windowWidth = window.screen.width;
      return Math.floor(windowWidth * 0.8);
    }
    return 500;
  }
}
