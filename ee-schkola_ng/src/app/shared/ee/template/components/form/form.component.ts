import {Component, Input, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';

@Component({
    selector: 'app-form',
    templateUrl: './form.component.html',
    styleUrls: ['./form.component.scss'],
    providers: [TemplateService],
})

export class FormComponent implements OnInit {

    @Input() formElement: any[][];

    constructor(public templateService: TemplateService) { }

    ngOnInit(): void {
        this.templateService.initElement(this.formElement);
    }
}
