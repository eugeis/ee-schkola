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
  index = 0;
  dataElement: any[][] = [];
  tempArray: any[][] = [];

  constructor() { }

  ngOnInit(): void {
  }
  inputElement() {
    this.graduation.name = this.simplifiedHtmlInputElement('graduationName');
    this.graduation.level = Number((<HTMLInputElement>document.getElementById('graduationLevel')).value);
    this.graduation.id = this.simplifiedHtmlInputElement('graduationId');
    this.other = this.simplifiedHtmlInputElement('other');
    this.profession = this.simplifiedHtmlInputElement('profession');
    this.dataElement.push([this.graduation.name, this.graduation.level, this.graduation.id, this.other, this.profession]);

  }
  deleteElement(index: number) {
    this.dataElement[index][0] = '';
    this.dataElement[index][1] = 0;
    this.dataElement[index][2] = '';
    this.dataElement[index][3] = '';
    this.dataElement[index][4] = '';
  }
  printElement(index: number) {
    console.log('GraduationName Value: ' + this.dataElement[index][0])
    console.log('GraduationLevel Value: ' + this.dataElement[index][1])
    console.log('GraduationId Value: ' + this.dataElement[index][2])
    console.log('Other Value: ' + this.dataElement[index][3])
    console.log('Profession Value: ' + this.dataElement[index][4])
  }
  changeIndex(input: number) {
        this.index = input;
  }
  loadElement(index: number) {
    (<HTMLInputElement>document.getElementById('graduationName')).value = this.dataElement[index][0];
    (<HTMLInputElement>document.getElementById('graduationLevel')).value = this.dataElement[index][1];
    (<HTMLInputElement>document.getElementById('graduationId')).value = this.dataElement[index][2];
    (<HTMLInputElement>document.getElementById('other')).value = this.dataElement[index][3];
    (<HTMLInputElement>document.getElementById('profession')).value = this.dataElement[index][4];
  }
  editElement(index: number) {
    this.dataElement[index][0] = this.simplifiedHtmlInputElement('graduationName');
    this.dataElement[index][1] = this.simplifiedHtmlInputElement('graduationLevel');
    this.dataElement[index][2] = this.simplifiedHtmlInputElement('graduationId');
    this.dataElement[index][3] = this.simplifiedHtmlInputElement('other');
    this.dataElement[index][4] = this.simplifiedHtmlInputElement('profession');
  }
  simplifiedHtmlInputElement(element: string) {
        return (<HTMLInputElement>document.getElementById(element)).value;
  }

}



