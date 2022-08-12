import {Component, OnInit} from '@angular/core';
import {ProfileDataService} from '../../services/profile-data.service';
import {TableDataService} from '../../../../template/services/data.service';
import {Profile} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-profile-list',
    templateUrl: './profile-list.component.html',
    styleUrls: ['./profile-list.component.scss'],
    providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileListComponent implements OnInit {

    profile: Profile = new Profile();

    tableHeader: Array<String> = [];

    constructor(public profileDataService: ProfileDataService) { }

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName',
            'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone',
            'email', 'cellphone', 'photo', 'maritalState', 'childrenCount', 'partner',
            'church', 'member', 'services', 'graduation', 'other', 'profession'];
    }
}
