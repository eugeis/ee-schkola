import {ExpensePurpose} from '@schkola/finance/FinanceApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ExpensePurposeDataService} from '../../service/expensepurpose-data.service';

@Component({
  selector: 'app-expensepurpose-view',
  templateUrl: './expensepurpose-entity-view.component.html',
  styleUrls: ['./expensepurpose-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpensePurposeDataService}]
})

export class ExpensePurposeViewComponent implements OnInit {


    expensepurpose: ExpensePurpose;

    constructor(@Inject(ExpensePurposeDataService) public expensepurposeDataService: ExpensePurposeDataService) {}

    ngOnInit(): void {
        this.expensepurpose = this.expensepurposeDataService.getFirst();
        
        this.expensepurposeDataService.checkRoute(this.expensepurpose);
    }
}




