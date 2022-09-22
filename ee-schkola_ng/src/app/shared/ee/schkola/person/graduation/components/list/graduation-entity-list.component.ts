import {Graduation} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {GraduationDataService} from '../../service/graduation-data.service';

@Component({
  selector: 'app-graduation-list',
  templateUrl: './graduation-entity-list.component.html',
  styleUrls: ['./graduation-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: GraduationDataService}]
})

export class GraduationListComponent implements OnInit {

    graduation: Graduation = new Graduation();

    tableHeader: Array<String> = [];

    constructor(@Inject(GraduationDataService) public graduationDataService: GraduationDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'level', 'id'];
    }
}




