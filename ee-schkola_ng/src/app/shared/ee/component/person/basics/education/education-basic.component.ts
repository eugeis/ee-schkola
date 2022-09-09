import {Education} from '../person/PersonApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.scss'],
})

export class EducationComponent implements OnInit {

    @Input() education: Education = new Education();

    ngOnInit(): void {
        if (this.education === undefined) {
            this.education = {graduation: {name: '', level: 0, id: ''}, other: '', profession: ''};
        }
    }
}




