import {Component, Input, OnInit} from '@angular/core';
import {FormService} from '../../../../template/services/form.service';

@Component({
    selector: 'app-person-name',
    templateUrl: './personname.component.html',
    styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent implements OnInit {

    @Input() usedService: FormService;

    constructor() { }

    ngOnInit(): void {

    }
}
