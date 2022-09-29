import {Education, Graduation} from '@schkola/person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-education',
  templateUrl: './education-basic.component.html',
  styleUrls: ['./education-basic.component.scss'],
})

export class EducationComponent implements OnInit {

    @Input() education: Education;
    
    ngOnInit() {
        if (this.education === undefined) {
            this.education = new Education();
        }
        if (this.education.graduation === undefined) {
            this.education.graduation = new Graduation();
        }
    }
    
}




