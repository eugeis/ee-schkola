import {Component, Input, OnInit} from '@angular/core';

@Component({
    selector: 'app-person-name',
    templateUrl: './personname.component.html',
    styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent implements OnInit {

    @Input() usedService: any;

    constructor() { }

    ngOnInit(): void {

    }
}
