import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ExpenseDataService} from '../../services/expense-data.service';

@Component({
  selector: 'app-expense-view',
  templateUrl: './expense-view.component.html',
  styleUrls: ['./expense-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpenseDataService}]
})

export class ExpenseViewComponent implements OnInit {


    expense: Expense;

    constructor(public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        this.expense = this.expenseDataService.getFirst();
        this.expenseDataService.checkRoute(this.expense);
    }
}




