import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ExpenseDataService} from '../../services/expense-data.service';

@Component({
  selector: 'app-expense-list',
  templateUrl: './expense-view.component.html',
  styleUrls: ['./expense-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ExpenseDataService}]
})

export class ExpenseViewComponent implements OnInit {

    expense: Expense = new Expense();

    tableHeader: Array<String> = [];

    constructor(public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'purpose', 'amount', 'profile', 'date', 'id'];
    }
}




