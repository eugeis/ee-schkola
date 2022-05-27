import { NgModule } from '@angular/core';
import {TemplateComponent} from './components/example/template.component';
import {ButtonComponent} from './components/button/button.component';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatOptionModule} from '@angular/material/core';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {ReactiveFormsModule} from '@angular/forms';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatGridListModule} from '@angular/material/grid-list';
import {FamilyComponent} from '../component/family/family.component';
import {EducationComponent} from '../component/education/education.component';
import {LocationComponent} from '../component/location/location.component';
import {PersonNameComponent} from '../component/personname/personname.component';
import {ContactComponent} from '../component/contact/contact.component';
import {ChurchInfoComponent} from '../component/churchinfo/churchinfo.component';
import {AddressComponent} from '../component/address/address.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {CommonModule} from '@angular/common';

@NgModule({
    declarations: [TemplateComponent,
        ButtonComponent,
        FamilyComponent,
        EducationComponent,
        LocationComponent,
        PersonNameComponent,
        ContactComponent,
        ChurchInfoComponent,
        AddressComponent,
    ],
    imports: [
        TemplateRoutingModules,
        CommonModule,
        MatInputModule,
        MatFormFieldModule,
        MatSelectModule,
        MatOptionModule,
        MatIconModule,
        MatButtonModule,
        ReactiveFormsModule,
        MatExpansionModule,
        MatGridListModule,
    ],
    providers: [

    ],
    exports: [
        TemplateComponent,
        ButtonComponent
    ]
})
export class TemplateModule { }
