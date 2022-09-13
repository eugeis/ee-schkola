import {Component, Input} from '@angular/core';
import {Address} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-address',
    templateUrl: './address.component.html',
    styleUrls: ['./address.component.scss']
})

export class AddressComponent {

    @Input() address: Address;
}
