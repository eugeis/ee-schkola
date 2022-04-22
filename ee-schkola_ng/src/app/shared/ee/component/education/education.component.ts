import { Component, OnInit } from '@angular/core';
import { Graduation } from '../../schkola/person/PersonApiBase';

@Component({
  selector: 'app-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.scss']
})

export class EducationComponent implements OnInit {

  graduation: Graduation = new Graduation()
  other: String
  profession: String

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.graduation.name = (<HTMLInputElement>document.getElementById('graduationName')).value;
    this.graduation.level = Number((<HTMLInputElement>document.getElementById('graduationLevel')).value);
    this.graduation.id = (<HTMLInputElement>document.getElementById('graduationId')).value;
    this.other = (<HTMLInputElement>document.getElementById('other')).value;
    this.profession = (<HTMLInputElement>document.getElementById('profession')).value;
  }
  deleteElement() {
    this.graduation.name = '';
    this.graduation.level = 0;
    this.graduation.id = '';
    this.other = '';
    this.profession = '';
  }
  printElement() {
    console.log('GraduationName Value: ' + this.graduation.name)
    console.log('GraduationLevel Value: ' + this.graduation.level)
    console.log('GraduationId Value: ' + this.graduation.id)
    console.log('Other Value: ' + this.other)
    console.log('Profession Value: ' + this.profession)
  }

}



