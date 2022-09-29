import {SchoolYear} from '@schkola/student/StudentApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {SchoolYearDataService} from '../../service/schoolyear-data.service';

@Component({
  selector: 'app-schoolyear-form',
  templateUrl: './schoolyear-form.component.html',
  styleUrls: ['./schoolyear-form.component.scss']
})

export class SchoolYearFormComponent implements OnInit {


    @Input() schoolyear: SchoolYear;

    constructor(public schoolyearDataService: SchoolYearDataService) {}

    ngOnInit(): void {
        
    }
}




