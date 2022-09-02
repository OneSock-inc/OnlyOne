import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/userService/user-service.service';
import { MessageBannerComponent } from 'src/app/message-banner/message-banner.component';
import { MessageBannerDirective as MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';

@Component({
  selector: 'app-abstract-form',
  template: ``,
  styles: []
})
export class AbstractFormComponent  {

  constructor(protected userService: UserService, protected router: Router) { }


  @ViewChild(MessageBannerDirective, {static: true})
  messageBanner!: MessageBannerDirective;

  

  displayError(message: string) {
    const elem = this.messageBanner.vcref.createComponent(MessageBannerComponent);
    elem.instance.message = message;
  }

  hideError() {
    this.messageBanner.vcref.clear();
  }

  

}

