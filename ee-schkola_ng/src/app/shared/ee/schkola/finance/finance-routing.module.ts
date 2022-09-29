import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {FinanceViewComponent} from './components/view/finance-module-view.component';
import {ExpenseViewComponent} from '@schkola/finance/expense/components/view/expense-entity-view.component';
import {ExpenseListComponent} from '@schkola/finance/expense/components/list/expense-entity-list.component';
import {ExpenseFormComponent} from '@schkola/finance/expense/components/form/expense-form.component';
import {ExpensePurposeViewComponent} from '@schkola/finance/expensepurpose/components/view/expensepurpose-entity-view.component';
import {ExpensePurposeListComponent} from '@schkola/finance/expensepurpose/components/list/expensepurpose-entity-list.component';
import {ExpensePurposeFormComponent} from '@schkola/finance/expensepurpose/components/form/expensepurpose-form.component';
import {FeeViewComponent} from '@schkola/finance/fee/components/view/fee-entity-view.component';
import {FeeListComponent} from '@schkola/finance/fee/components/list/fee-entity-list.component';
import {FeeFormComponent} from '@schkola/finance/fee/components/form/fee-form.component';
import {FeeKindViewComponent} from '@schkola/finance/feekind/components/view/feekind-entity-view.component';
import {FeeKindListComponent} from '@schkola/finance/feekind/components/list/feekind-entity-list.component';
import {FeeKindFormComponent} from '@schkola/finance/feekind/components/form/feekind-form.component';

const routes: Routes = [
    { path: '', component: FinanceViewComponent },
    { path: 'expense', component: ExpenseListComponent },
    { path: 'expense/new', component: ExpenseViewComponent },
    { path: 'expense/edit/:id', component: ExpenseViewComponent },
    { path: 'expensepurpose', component: ExpensePurposeListComponent },
    { path: 'expensepurpose/new', component: ExpensePurposeViewComponent },
    { path: 'expensepurpose/edit/:id', component: ExpensePurposeViewComponent },
    { path: 'fee', component: FeeListComponent },
    { path: 'fee/new', component: FeeViewComponent },
    { path: 'fee/edit/:id', component: FeeViewComponent },
    { path: 'feekind', component: FeeKindListComponent },
    { path: 'feekind/new', component: FeeKindViewComponent },
    { path: 'feekind/edit/:id', component: FeeKindViewComponent }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class FinanceRoutingModules {}





