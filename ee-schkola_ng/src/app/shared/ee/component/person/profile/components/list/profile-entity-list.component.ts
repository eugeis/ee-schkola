import {Profile} from '../person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ProfileDataService} from '../../services/profile-data.service';

@Component({
  selector: 'app-profile-list',
  templateUrl: './profile-view.component.html',
  styleUrls: ['./profile-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    profile: Profile = new Profile();

    tableHeader: Array<String> = [];

    constructor(public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'name', 'birthname', 'birthday', 'address', 'contact', 'photodata', 'photo', 'family', 'church', 'education', 'id'];
    }
}




