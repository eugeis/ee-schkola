import {Component, Input, OnInit} from '@angular/core';
import {Address} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-address',
    templateUrl: './address.component.html',
    styleUrls: ['./address.component.scss']
})

export class AddressComponent implements OnInit {

    @Input() address: Address = new Address();

    constructor() { }

    ngOnInit(): void {
        if (this.address === undefined) {
            this.address = {code: '', city: '', country: '', suite: '', street: ''};
        }
    }
}
