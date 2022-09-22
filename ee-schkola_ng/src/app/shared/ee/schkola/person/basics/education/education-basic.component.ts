import {Education} from '@schkola/person/PersonApiBase'

import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-education',
  templateUrl: './education-basic.component.html',
  styleUrls: ['./education-basic.component.scss'],
})

export class EducationComponent {

    @Input() education: Education;
    
}




