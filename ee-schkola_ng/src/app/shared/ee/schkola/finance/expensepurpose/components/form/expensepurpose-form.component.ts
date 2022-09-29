import {ExpensePurpose} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {ExpensePurposeDataService} from '../../service/expensepurpose-data.service';

@Component({
  selector: 'app-expensepurpose-form',
  templateUrl: './expensepurpose-form.component.html',
  styleUrls: ['./expensepurpose-form.component.scss']
})

export class ExpensePurposeFormComponent implements OnInit {


    @Input() expensepurpose: ExpensePurpose;

    constructor(public expensepurposeDataService: ExpensePurposeDataService) {}

    ngOnInit(): void {
        
    }
}




