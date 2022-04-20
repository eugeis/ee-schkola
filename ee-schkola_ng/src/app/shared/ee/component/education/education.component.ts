import { Component, OnInit } from '@angular/core';
import { Graduation } from '../../schkola/person/PersonApiBase';

@Component({
  selector: 'app-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.css']
})

export class EducationComponent implements OnInit {

  graduation: Graduation
  other: string
  profession: string

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.graduation = (<HTMLInputElement>document.getElementById('graduation')).value;
    this.other = (<HTMLInputElement>document.getElementById('other')).value;
    this.profession = (<HTMLInputElement>document.getElementById('profession')).value;
  }
  deleteElement() {
    this.graduation = new Graduation();
    this.other = '';
    this.profession = '';
  }
  printElement() {
    console.log('Graduation Value: ' + this.graduation)
    console.log('Other Value: ' + this.other)
    console.log('Profession Value: ' + this.profession)
  }

}



