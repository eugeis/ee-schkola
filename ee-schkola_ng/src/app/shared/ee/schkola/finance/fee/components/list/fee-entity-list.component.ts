import {Fee} from '@schkola/finance/FinanceApiBase'

import {Component, Inject, OnInit} from '@angular/core';
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

    constructor(@Inject(FeeDataService) public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'amount', 'kind', 'date', 'id'];
    }
}




