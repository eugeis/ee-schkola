import {Contact, PersonName, Profile} from '@schkola/person/PersonApiBase'
import {SchoolApplication, SchoolYear} from '@schkola/student/StudentApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {SchoolApplicationDataService} from '../../service/schoolapplication-data.service';

@Component({
  selector: 'app-schoolapplication-form',
  templateUrl: './schoolapplication-form.component.html',
  styleUrls: ['./schoolapplication-form.component.scss']
})

export class SchoolApplicationFormComponent implements OnInit {


    @Input() schoolapplication: SchoolApplication;

    constructor(public schoolapplicationDataService: SchoolApplicationDataService) {}

    ngOnInit(): void {
        if (this.schoolapplication.profile === undefined) {
            this.schoolapplication.profile = new Profile();
        }
        if (this.schoolapplication.churchContactPerson === undefined) {
            this.schoolapplication.churchContactPerson = new PersonName();
        }
        if (this.schoolapplication.churchContact === undefined) {
            this.schoolapplication.churchContact = new Contact();
        }
        if (this.schoolapplication.schoolYear === undefined) {
            this.schoolapplication.schoolYear = new SchoolYear();
        }
    }
}




