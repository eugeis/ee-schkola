import {Fee} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {FeeDataService} from '../../service/fee-data.service';

@Component({
  selector: 'app-fee-list',
  templateUrl: './fee-entity-list.component.html',
  styleUrls: ['./fee-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeDataService}]
})

export class FeeListComponent implements OnInit {

    fee: Fee = new Fee();

    tableHeader: Array<String> = [];

    constructor(public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'amount', 'name', 'amount', 'description', 'id', 'date', 'id'];
    }
}




