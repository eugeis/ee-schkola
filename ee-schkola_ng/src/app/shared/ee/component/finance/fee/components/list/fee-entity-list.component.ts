import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {FeeDataService} from '../../services/fee-data.service';

@Component({
  selector: 'app-fee-list',
  templateUrl: './fee-view.component.html',
  styleUrls: ['./fee-view.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeDataService}]
})

export class FeeViewComponent implements OnInit {

    fee: Fee = new Fee();

    tableHeader: Array<String> = [];

    constructor(public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'student', 'amount', 'kind', 'date', 'id'];
    }
}




