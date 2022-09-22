import {Expense} from '@schkola/finance/FinanceApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ExpenseDataService} from '../../service/expense-data.service';

@Component({
  selector: 'app-expense-view',
  templateUrl: './expense-entity-view.component.html',
  styleUrls: ['./expense-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpenseDataService}]
})

export class ExpenseViewComponent implements OnInit {


    expense: Expense;

    constructor(@Inject(ExpenseDataService) public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        this.expense = this.expenseDataService.getFirst();
        
        this.expenseDataService.checkRoute(this.expense);
    }
}




