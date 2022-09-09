import {Church} from '../person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ChurchDataService} from '../../services/church-data.service';

@Component({
  selector: 'app-church-list',
  templateUrl: './church-view.component.html',
  styleUrls: ['./church-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchViewComponent implements OnInit {

    church: Church = new Church();

    tableHeader: Array<String> = [];

    constructor(public churchDataService: ChurchDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'address', 'pastor', 'contact', 'association', 'id'];
    }
}




