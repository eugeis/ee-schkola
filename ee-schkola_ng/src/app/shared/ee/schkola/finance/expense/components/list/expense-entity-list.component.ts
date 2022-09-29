import {Expense} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit} from '@angular/core';
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

    constructor(public expenseDataService: ExpenseDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'description', 'id', 'amount', 'gender', 'first', 'last', 'birthName', 'birthday', 'street', 'suite', 'city', 'code', 'country', 'phone', 'email', 'cellphone', 'photoData', 'photo', 'maritalState', 'childrenCount', 'first', 'last', 'church', 'member', 'services', 'name', 'level', 'id', 'other', 'profession', 'id', 'date', 'id'];
    }
}




