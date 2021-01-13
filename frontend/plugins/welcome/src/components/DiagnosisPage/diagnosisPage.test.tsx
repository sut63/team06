import React from 'react';
import { render } from '@testing-library/react';
import DiagnosisPage from './diagnosisPage';
import { ThemeProvider } from '@material-ui/core';
import { lightTheme } from '@backstage/theme';

describe('DiagnosisPage', () => {
    it('should render', () => {
        const rendered = render(
            <ThemeProvider theme={lightTheme}>
                <DiagnosisPage />
            </ThemeProvider>,
        );
        expect(rendered.baseElement).toBeInTheDocument();
    });
});
