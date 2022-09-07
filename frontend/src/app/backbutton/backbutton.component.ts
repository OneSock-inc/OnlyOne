import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-backbutton',
  templateUrl: './backbutton.component.html',
  styleUrls: ['./backbutton.component.scss']
})
export class BackbuttonComponent implements OnInit {
  @Input() page = "/";

  constructor() { }

  ngOnInit(): void {
  }
}
