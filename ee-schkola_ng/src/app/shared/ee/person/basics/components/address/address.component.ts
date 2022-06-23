import {Component, Input, OnInit} from '@angular/core';
import {FormService} from '../../../../template/services/form.service';

@Component({
    selector: 'app-address',
    templateUrl: './address.component.html',
    styleUrls: ['./address.component.scss']
})

export class AddressComponent implements OnInit {

    @Input() usedService: FormService;

    constructor() { }

    ngOnInit(): void {

    }
}
