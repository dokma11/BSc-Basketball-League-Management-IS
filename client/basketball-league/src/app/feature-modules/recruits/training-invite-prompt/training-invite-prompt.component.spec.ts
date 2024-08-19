import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TrainingInvitePromptComponent } from './training-invite-prompt.component';

describe('TrainingInvitePromptComponent', () => {
  let component: TrainingInvitePromptComponent;
  let fixture: ComponentFixture<TrainingInvitePromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TrainingInvitePromptComponent]
    });
    fixture = TestBed.createComponent(TrainingInvitePromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
