import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ExpensePurposeDataService} from '../../services/expensepurpose-data.service';

@Component({
  selector: 'app-expensepurpose-view',
  templateUrl: './expensepurpose-view.component.html',
  styleUrls: ['./expensepurpose-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpensePurposeDataService}]
})

export class ExpensePurposeViewComponent implements OnInit {


    expensepurpose: ExpensePurpose;

    constructor(public expensepurposeDataService: ExpensePurposeDataService) {}

    ngOnInit(): void {
        this.expensepurpose = this.expensepurposeDataService.getFirst();
        this.expensepurposeDataService.checkRoute(this.expensepurpose);
    }
}




