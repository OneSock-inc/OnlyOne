import { Component, OnInit } from '@angular/core';

type LinkElement = {
  text: string;
  href: string;
  classlist: string;
};

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {

  constructor() { }

  linksClassList: string = 'center btn';

  links: Array<LinkElement> = [
    {text: "login", href: '/login', classlist: this.linksClassList},
    {text: "Add a lonely sock", href: '/add-sock', classlist: this.linksClassList},
    {text: "My socks", href: '/my-socks', classlist: this.linksClassList}
  ]

  ngOnInit(): void {
  }

}