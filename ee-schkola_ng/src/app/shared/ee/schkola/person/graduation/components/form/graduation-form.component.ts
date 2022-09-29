import {Graduation, GraduationLevel} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {GraduationDataService} from '../../service/graduation-data.service';

@Component({
  selector: 'app-graduation-form',
  templateUrl: './graduation-form.component.html',
  styleUrls: ['./graduation-form.component.scss']
})

export class GraduationFormComponent implements OnInit {

    graduationlevelEnum = this.graduationDataService.loadEnumElement(GraduationLevel);
    @Input() graduation: Graduation;

    constructor(public graduationDataService: GraduationDataService) {}

    ngOnInit(): void {
        
    }
}




