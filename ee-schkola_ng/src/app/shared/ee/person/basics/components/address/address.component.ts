import {Component, Input, OnInit} from '@angular/core';
import {FormService} from '../../../../template/services/form.service';
import {Address} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-address',
    templateUrl: './address.component.html',
    styleUrls: ['./address.component.scss']
})

export class AddressComponent implements OnInit {

    @Input() usedService: FormService;
    @Input() address: Address;

    constructor() { }

    ngOnInit(): void {

    }
}
