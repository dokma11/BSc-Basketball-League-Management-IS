import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeclineRequestPromptComponent } from './decline-request-prompt.component';

describe('DeclineRequestPromptComponent', () => {
  let component: DeclineRequestPromptComponent;
  let fixture: ComponentFixture<DeclineRequestPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DeclineRequestPromptComponent]
    });
    fixture = TestBed.createComponent(DeclineRequestPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
