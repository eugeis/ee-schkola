import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ExpensePurposeDataService} from '../../services/expensepurpose-data.service';

@Component({
  selector: 'app-expensepurpose-list',
  templateUrl: './expensepurpose-view.component.html',
  styleUrls: ['./expensepurpose-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpensePurposeDataService}]
})

export class ExpensePurposeViewComponent implements OnInit {

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




