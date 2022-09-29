import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ProfileDataService} from '../../service/profile-data.service';

@Component({
  selector: 'app-profile-list',
  templateUrl: './profile-entity-list.component.html',
  styleUrls: ['./profile-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileListComponent implements OnInit {

    profile: Profile = new Profile();

    tableHeader: Array<String> = [];

    constructor(public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id'];
    }
}




