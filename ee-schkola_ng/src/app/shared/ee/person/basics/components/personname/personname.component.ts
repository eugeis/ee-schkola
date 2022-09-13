import {Component, Input} from '@angular/core';
import {PersonName} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-name',
    templateUrl: './personname.component.html',
    styleUrls: ['./personname.component.scss']
})

export class PersonNameComponent {

    @Input() personName: PersonName;
}
