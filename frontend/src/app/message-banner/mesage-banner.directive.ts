import { Directive, ViewContainerRef } from '@angular/core';

@Directive({
  selector: '[appMessageBanner]'
})
export class MesageBannerDirective {

  constructor(public vcref: ViewContainerRef) { }

}
