import {Expense} from '@schkola/finance/FinanceApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ExpenseDataService} from '../../service/expense-data.service';

@Component({
  selector: 'app-expense-list',
  templateUrl: './expense-entity-list.component.html',
  styleUrls: ['./expense-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpenseDataService}]
})

export class ExpenseListComponent implements OnInit {

    expense: Expense = new Expense();

    tableHeader: Array<String> = [];

    constructor(@Inject(ExpenseDataService) public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'purpose', 'amount', 'profile', 'date', 'id'];
    }
}




