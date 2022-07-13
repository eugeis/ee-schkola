import {Component, Input, OnInit} from '@angular/core';
import {FormService} from '../../../../template/services/form.service';
import {PersonName} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-name',
    templateUrl: './personname.component.html',
    styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent implements OnInit {

    @Input() usedService: FormService;
    @Input() personName: PersonName;

    constructor() { }

    ngOnInit(): void {

    }
}
