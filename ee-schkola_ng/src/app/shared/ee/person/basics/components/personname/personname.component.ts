import {Component, Input, OnInit} from '@angular/core';
import {PersonName} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-name',
    templateUrl: './personname.component.html',
    styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent implements OnInit {

    @Input() personName: PersonName;

    constructor() { }

    ngOnInit(): void {

    }
}
