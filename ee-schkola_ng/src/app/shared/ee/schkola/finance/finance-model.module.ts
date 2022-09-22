import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {FinanceRoutingModules} from './finance-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {FinanceViewComponent} from './components/view/finance-module-view.component';
import {ExpenseViewComponent} from './expense/components/view/expense-entity-view.component';
import {ExpenseListComponent} from './expense/components/list/expense-entity-list.component';
import {ExpensePurposeViewComponent} from './expensepurpose/components/view/expensepurpose-entity-view.component';
import {ExpensePurposeListComponent} from './expensepurpose/components/list/expensepurpose-entity-list.component';
import {FeeViewComponent} from './fee/components/view/fee-entity-view.component';
import {FeeListComponent} from './fee/components/list/fee-entity-list.component';
import {FeeKindViewComponent} from './feekind/components/view/feekind-entity-view.component';
import {FeeKindListComponent} from './feekind/components/list/feekind-entity-list.component';
import {PersonModule} from '@schkola/person/person-model.module';



@NgModule({
    declarations: [
        FinanceViewComponent,
        ExpenseViewComponent,
        ExpenseListComponent,
        ExpensePurposeViewComponent,
        ExpensePurposeListComponent,
        FeeViewComponent,
        FeeListComponent,
        FeeKindViewComponent,
        FeeKindListComponent,

    ],
    imports: [
        FinanceRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule,
        PersonModule
    ],
    providers: [],
    exports: [
        ExpenseViewComponent,
        ExpensePurposeViewComponent,
        FeeViewComponent,
        FeeKindViewComponent,

    ]
})
export class FinanceModule { }




