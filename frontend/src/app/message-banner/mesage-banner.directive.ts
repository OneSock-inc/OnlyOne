import { Directive, ViewContainerRef } from '@angular/core';

@Directive({
  selector: '[appMesageBanner]'
})
export class MesageBannerDirective {

  constructor(public vcref: ViewContainerRef) { }

}
