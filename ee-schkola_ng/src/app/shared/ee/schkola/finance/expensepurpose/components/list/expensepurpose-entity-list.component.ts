import {ExpensePurpose} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ExpensePurposeDataService} from '../../service/expensepurpose-data.service';

@Component({
  selector: 'app-expensepurpose-list',
  templateUrl: './expensepurpose-entity-list.component.html',
  styleUrls: ['./expensepurpose-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpensePurposeDataService}]
})

export class ExpensePurposeListComponent implements OnInit {

    expensepurpose: ExpensePurpose = new ExpensePurpose();

    tableHeader: Array<String> = [];

    constructor(public expensepurposeDataService: ExpensePurposeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'description', 'id'];
    }
}




