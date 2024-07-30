import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CancelRequestPromptComponent } from './cancel-request-prompt.component';

describe('CancelRequestPromptComponent', () => {
  let component: CancelRequestPromptComponent;
  let fixture: ComponentFixture<CancelRequestPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CancelRequestPromptComponent]
    });
    fixture = TestBed.createComponent(CancelRequestPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
