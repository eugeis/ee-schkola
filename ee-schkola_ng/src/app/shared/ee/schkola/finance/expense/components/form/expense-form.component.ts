import {Expense, ExpensePurpose} from '@schkola/finance/FinanceApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {ExpenseDataService} from '../../service/expense-data.service';

@Component({
  selector: 'app-expense-form',
  templateUrl: './expense-form.component.html',
  styleUrls: ['./expense-form.component.scss']
})

export class ExpenseFormComponent implements OnInit {


    @Input() expense: Expense;

    constructor(public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        if (this.expense.purpose === undefined) {
            this.expense.purpose = new ExpensePurpose();
        }
        if (this.expense.profile === undefined) {
            this.expense.profile = new Profile();
        }
    }
}




