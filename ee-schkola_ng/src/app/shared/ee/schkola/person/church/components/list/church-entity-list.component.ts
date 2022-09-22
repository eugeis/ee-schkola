import {Church} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ChurchDataService} from '../../service/church-data.service';

@Component({
  selector: 'app-church-list',
  templateUrl: './church-entity-list.component.html',
  styleUrls: ['./church-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchListComponent implements OnInit {

    church: Church = new Church();

    tableHeader: Array<String> = [];

    constructor(@Inject(ChurchDataService) public churchDataService: ChurchDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'address', 'pastor', 'contact', 'association', 'id'];
    }
}




