import {Address, Church, Contact, PersonName} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {ChurchDataService} from '../../service/church-data.service';

@Component({
  selector: 'app-church-form',
  templateUrl: './church-form.component.html',
  styleUrls: ['./church-form.component.scss']
})

export class ChurchFormComponent implements OnInit {


    @Input() church: Church;

    constructor(public churchDataService: ChurchDataService) {}

    ngOnInit(): void {
        if (this.church.address === undefined) {
            this.church.address = new Address();
        }
        if (this.church.pastor === undefined) {
            this.church.pastor = new PersonName();
        }
        if (this.church.contact === undefined) {
            this.church.contact = new Contact();
        }
    }
}




